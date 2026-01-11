import { TAFLog } from "./fee/taf-fee";
import { getStockParser } from "./stock/factory";
import { getDividendParser } from "./dividend/factory";
import type { Parser } from "./parser";
import type { Dividend } from "./dividend/dividend";
import type { Stock } from "./stock/stock";
import type { Fee } from "./fee/fee";

export function getParser(text: string): Parser<Stock | Dividend | Fee> {
  if (["Sell", "Buy"].some((s) => text.startsWith(s)))
    return getStockParser(text);
  if (["Dividend"].some((s) => text.startsWith(s)))
    return getDividendParser(text);
  if (text.startsWith("TAF")) return new TAFLog(text);
  throw new Error("no parser found");
}
export function saveLog<T>(parser: Parser<T>) {
  parser.save();
}
