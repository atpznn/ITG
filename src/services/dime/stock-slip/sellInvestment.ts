import type { BasePatternExtractor } from "../../extracter/patterns/base-pattern-extractor";
import {
  dateRegexWithOutPMAM,
  extractDateFromText,
  extractDatesFromText,
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

function getBefore(text: string, keyword: string, count: number) {
  const pattern = `(\\S+)\\s+`.repeat(count) + `${keyword}`;
  const match = text.match(new RegExp(pattern, "i"));
  if (!match) throw new Error(`not found ${keyword}`);
  return match[0];
}
export class SellInvestmentLog implements IInvestmentLog {
  constructor(private words: string, private extractor: BasePatternExtractor) {}
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
    const price = getFloat(getWordAfter(orderDetail, "Total Credit", 2));
    const executedPrice = getFloat(getWordAfter(orderDetail, "Executed", 4));
    const shares = getFloat(getBefore(orderDetail, "Shares", 1));
    const stockAmount = getFloat(getBefore(orderDetail, "Amount", 3));
    const vat: Vat = {
      commissionFee: getTax(orderDetail, "Commission"),
      secFee: getTax(orderDetail, "SEC"),
      tafFee: getTax(orderDetail, "TAF"),
      vat7: getTax(orderDetail, "VAT 7%"),
    };
    const allVatPrice = sumVat(vat);
    const diffPrice = stockAmount - price;
    const diffVat = diffPrice - allVatPrice;
    return {
      kind: "slip",
      type: getType(this.words) as InvestmentType,
      symbol: getSymbol(this.words)!,
      stockAmount,
      allVatPrice,
      executedPrice,
      completionDate,
      diffVat,
      vat,
      vatExecuted: diffPrice,
      value: price,
      shares,
      submissionDate,
    };
  }
}
