import fs from "fs";
import {
  parseImageToText,
  readImageBufferFromPath,
} from "./services/ocr/index";
import { createAInvestmentLog } from "./services/dime/stock-slip/core";
import { CoordinatesOcrStategy } from "./services/ocr/stategies/coordinates-ocr";

// fs.readdir("./imageTest/dime", async (err, langFolders) => {
//   for (const langFolder of langFolders) {
//     const folders = fs.readdirSync(`./imageTest/dime/${langFolder}`);
//     for (const folder of folders) {
//       const files = fs.readdirSync(`./imageTest/dime/${langFolder}/${folder}`);

//       for (const file of files) {
//         const text = await parseImageToText(
//           await readImageBufferFromPath(
//             `./imageTest/dime/${langFolder}/${folder}/${file}`
//           )
//         );
//         console.log(
//           `${langFolder}/${folder}/${file}: \n\n${text}\n---------------------\n`
//         );
//         // const dateExtractor = new DatePatternExtractor();
//         // const transaction = new TransactionExtractor(dateExtractor, text);
//         // console.log(transaction.toJson());
//       }
//     }
//   }
// });
const basePath = "./imageTest/dime/en/slips";
fs.readdir(basePath, async (err, files) => {
  if (err) return console.log(err);
  for (const file of files) {
    const text = await parseImageToText(
      await readImageBufferFromPath(`${basePath}/${file}`),
      new CoordinatesOcrStategy()
    );
    console.log(createAInvestmentLog(text).toJson());
  }
});

// console.log(text)
// for (const element of [1, 1, 1, 1, 1, 1, 1, 1]) {
//   const text = await parseImageToText(
//     await readImageBufferFromPath(
//       `imageTest/1000001054.jpg`
//     ), new CoordinatesOcrStategy()
//   );
//   const h = ((await (new TransactionExtractor(new DatePatternExtractor(), text).toJson())))
//   const k = ((await createAInvestmentLog(await parseImageToText(
//     await readImageBufferFromPath(
//       `imageTest/1000000979.jpg`
//     ), new CoordinatesOcrStategy())).toJson()))
//   // console.log(h, k)
// }

// const paragraphs = `
// 1+2  1+4  1+5
// `;
// const p1 = /(\+)/g;
// const p2 = /\+/g;
// const p3 = /(?<=\+)/g;
// const p4 = /(?=\+)/g;
// const p5 = /(?:\+)/g;
// const f = "Completion date 28 Mar 2024 - 22:13";
// console.log(extractDateFromText(f, dateRegexWithOutPMAM));
// console.log(paragraphs.split(p1));
// console.log(paragraphs.split(p2));
// console.log(paragraphs.split(p3));
// console.log(paragraphs.split(p4));
// console.log(paragraphs.split(p5));

interface a {
  a: number;
  b: number;
  c: number;
}
function blockPlus(input1: number, input2: number) {
  return input1 + input2;
}
function blockMinus(input1: number, input2: number) {
  return input2 - input1;
}

// const as: a[] = [
//   { a: 1, b: 2, c: 3 },
//   { a: 1, b: 11, c: 7 },
//   { a: 10, b: 23, c: 13 },
//   { a: 1, b: 2, c: 3 },
// ];
// const s = as.map((s) => blockMinus(s["a"], s["b"]) + s.c);
// console.log(s);
