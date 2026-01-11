import tesseract from "node-tesseract-ocr";
import fs from "fs";
import path from 'path'
import sharp, { type Sharp } from 'sharp';
import type { OcrStategy } from "./stategies/base-ocr";
export async function readImageBufferFromPath(path: string) {
  const image = await fs.readFileSync(path);
  const imageBuffer = Buffer.from(image.buffer);
  return imageBuffer;
}
function sanitize(text: string): string {
  return text.replace(/([ก-์])\s+(?=[ก-์])/g, "$1").replace(/\s+/g, " ");
}
async function saveImage(inputBuffer: Buffer, outputFileName: string) {
  try {
    const debugDir = './imageTest/debugImages';
    if (!fs.existsSync(debugDir)) {
      fs.mkdirSync(debugDir, { recursive: true });
    }
    const outputPath = path.join(debugDir, outputFileName);
    await fs.promises.writeFile(outputPath, inputBuffer);
    console.log(`Successfully saved: ${outputPath}`);

  } catch (error) {
    console.error('Error processing image:', error);
    throw error;
  }
}
async function getIsDark(buffer: Buffer) {
  // สร้างภาพขนาดจิ๋ว (เช่น 10x10 px) เพื่อหาค่าเฉลี่ยสีเท่านั้น
  const tinyStats = await sharp(buffer).resize(10, 10).stats();
  const mean = tinyStats.channels.reduce((acc, c) => acc + c.mean, 0) / tinyStats.channels.length;
  return mean < 128;
}
async function normalizeWithTheme(sharpImage: sharp.Sharp) {
  const stats = await sharpImage.stats();
  const isDark = stats.channels.reduce((acc, c) => acc + c.mean, 0) / stats.channels.length < 128;
  const threshold = isDark ? 120 : 200
  if (isDark) return { image: sharpImage.modulate({ brightness: 1.8 }), threshold }
  return {
    image: sharpImage.modulate({
      brightness: 0.3,
      saturation: 0
    }), threshold
  }

}
async function fastResize(buffer: Buffer, percent: number = 50) {
  const image = sharp(buffer);

  // 1. ดึงขนาดต้นฉบับ (metadata ใช้เวลาน้อยมาก ไม่เหมือน stats)
  const meta = await image.metadata();

  if (meta.width && meta.height) {
    const newWidth = Math.round(meta.width * (percent / 100));

    return await image
      .resize(newWidth) // ใส่แค่ Width เดี๋ยว Sharp จัดการ Height ให้เองตามสัดส่วน
      .grayscale()      // ช่วยลดขนาด Buffer ลงได้อีก 3 เท่า (จาก 3 channels เหลือ 1)
      .toBuffer();
  }

  return buffer; // ถ้าดึงขนาดไม่ได้ ให้ส่งตัวเดิมกลับ
}
function convertBackToWrite(image: Sharp, isDark: boolean) {
  if (isDark) return image.negate({ alpha: false })
  // .modulate({ brightness: isDark ? 2 : 0.4 })
  return image
}
async function greyScale(buffer: Buffer) {
  sharp.concurrency(1)
  sharp.cache(false);
  // const sharpImage = sharp(buffer);
  const isDark = await getIsDark(buffer);
  const threshold = isDark ? 120 : 200;
  // const { image, threshold } = await normalizeWithTheme(
  // sharpImage
  // .resize(1600, null, { withoutEnlargement: true })
  // .withMetadata({ density: 300 })
  // .ensureAlpha()
  // )
  return convertBackToWrite(
    await sharp(await fastResize(buffer, 60))
    // .ensureAlpha()
    // .modulate({ brightness: isDark ? 2 : 0.4 })
    , isDark
  )
    // .greyscale()
    // .threshold(threshold)
    // .png({
    //   quality: 100,
    //   compressionLevel: 9 // บีบอัดขนาดไฟล์ให้เล็กที่สุด (ประหยัด RAM ตอนส่งให้ OCR)
    // })
    .toBuffer();
}
export function cleanText(text: string) {
  return sanitize(text)
    .split("\n")
    .map((line) => line.trim())
    .filter((line) => line !== "")
    .join(" ");
}
export async function parseImageToText(image: Buffer, ocrStategy: OcrStategy) {
  // console.time("Image-Processing");
  // const greyImage = await greyScale(image)
  // console.timeEnd("Image-Processing");
  // await saveImage(greyImage, `ocr-ready-${Date.now()}.png`)
  // console.time("OCR-Task only");
  // await tesseract.recognize(image, ocrStategy.getConfig())
  // console.timeEnd("OCR-Task only");
  console.time("OCR-Task with preprocess");
  const tsvData = await tesseract.recognize(image, ocrStategy.getConfig())
  console.timeEnd("OCR-Task with preprocess");
  const text = ocrStategy.mutation(tsvData)
  return cleanText(text)
}
