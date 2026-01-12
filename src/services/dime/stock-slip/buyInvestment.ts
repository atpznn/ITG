import type { BasePatternExtractor } from "../../extracter/patterns/base-pattern-extractor";
import {
  dateRegexWithOutPMAM,
  extractDateFromText,
  extractDatesFromText,
  findSame,
  findWord,
  findWordUseNextLine,
  parseDateTimeToDateObject,
  parseUsd,
} from "../../util";
import type { IInvestmentLog, Investment, InvestmentType, Vat } from "./core";
import { getSymbol, getType, sumVat } from "./util";

function getTax(text: string, keyword: string) {
  try {
    const tax = getWordAfter(text, keyword, 3);
    return getFloat(tax.replace(keyword, ""));
  } catch (ex) {
    return 0;
  }
}
function getWordAfter(text: string, keyword: string, count: number) {
  const pattern = `${keyword}` + `\\s+(\\S+)`.repeat(count);
  const regex = new RegExp(pattern, "i");
  const match = text.match(regex);
  if (!match) throw new Error(`not found ${keyword}`);
  return match[0];
}
function getFloat(text: string) {
  const regex = new RegExp(/-?\d+\.?\d*/g);
  const match = text.match(regex);
  if (!match) throw new Error(`not found number`);
  return parseFloat(match[0]);
}
function tryGetFloat(text: string) {
  try {
    return getFloat(text);
  } catch (ex) {
    return undefined;
  }
}
function getBefore(text: string, keyword: string, count: number) {
  const pattern = `(\\S+)\\s+`.repeat(count) + `${keyword}`;
  const match = text.match(new RegExp(pattern, "i"));
  if (!match) throw new Error(`not found ${keyword}`);
  return match[0];
}
export class BuyInvestmentLog implements IInvestmentLog {
  constructor(private words: string, private extractor: BasePatternExtractor) {
    // console.log(words);
  }
  toJson(): Investment {
    const [orderDetail, dateDetail] = this.words
      .replace(/[\u0E00-\u0E7F]/g, "")
      .split(/(?=\Submission Date)/g);
    if (!orderDetail) throw new Error("no order detail data");
    if (!dateDetail) throw new Error("no date data");
    const [submissionDate, completionDate] = this.extractor
      .extract(dateDetail)
      .map((x) => extractDateFromText(x, dateRegexWithOutPMAM));
    if (!submissionDate) throw new Error("no submission date");
    if (!completionDate) throw new Error("no completion date");
    const [executedPrice, shares] = getWordAfter(orderDetail, "Shares", 8)
      .split(" ")
      .map(tryGetFloat)
      .filter((x) => x != undefined);
    if (!executedPrice) throw new Error("not found Execute Price");
    if (!shares) throw new Error("not found shares");
    const stockAmount = getFloat(getBefore(orderDetail, "Amount", 3));
    const vat: Vat = {
      commissionFee: getTax(orderDetail, "Commission"),
      secFee: getTax(orderDetail, "SEC"),
      tafFee: getTax(orderDetail, "TAF"),
      vat7: getTax(orderDetail, "VAT 7%"),
    };
    const allVatPrice = sumVat(vat);
    const diffPrice = 0;
    const diffVat = diffPrice - allVatPrice;
    return {
      kind: "slip",
      type: getType(this.words) as InvestmentType,
      symbol: getSymbol(this.words)!,
      stockAmount: stockAmount,
      allVatPrice,
      executedPrice,
      completionDate,
      vat,
      vatExecuted: diffPrice,
      diffVat,
      value: stockAmount,
      shares,
      submissionDate,
    };
  }
}
