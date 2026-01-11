/** @type {import('ts-jest').JestConfigWithTsJest} */
export default {
  preset: "ts-jest/presets/default-esm", // สำคัญมาก: ใช้ Preset สำหรับ ESM
  testEnvironment: "node",
  extensionsToTreatAsEsm: [".ts"], // บอกให้ Jest ปฏิบัติกับ .ts แบบ ESM
  moduleNameMapper: {
    "^(\\.{1,2}/.*)\\.js$": "$1", // จัดการเรื่องการ import ที่มีนามสกุล .js ในไฟล์ TS
  },
  transform: {
    "^.+\\.tsx?$": [
      "ts-jest",
      {
        useESM: true, // เปิดการใช้งาน ESM ในตัวแปลงไฟล์
      },
    ],
  },
};
