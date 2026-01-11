
import type { Express } from "express";
import imageProcessRoute from "./image-process";
import { createAInvestmentLog } from "../../services/dime/stock-slip/core";
import { DatePatternExtractor } from "../../services/extracter/patterns/date-pattern-extractor";
import { TransactionExtractor } from "../../services/dime/transaction/transaction-extractor";
import { CoordinatesOcrStategy } from "../../services/ocr/stategies/coordinates-ocr";
import { cleanText } from "../../services/ocr";
const dimeHandler = (text: string) => {
  try {
    if (text.includes('Stock Amount')) {
      return { slip: [createAInvestmentLog(text).toJson()] }
    }
    const dateExtractor = new DatePatternExtractor();
    const extractor = new TransactionExtractor(dateExtractor, text);
    return { ...extractor.toJson() }
  }
  catch (ex) {
    throw ex
  }
};
export default function v1DimeRoute(app: Express) {
  app.use("/v1/dime", imageProcessRoute(dimeHandler));
  app.post('/v1/dime/process-text', (req, res) => {
    const text = cleanText(req.body.text)
    console.log(text)
    res.json(dimeHandler(text))
  })
  return app
}
