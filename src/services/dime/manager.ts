import { DatePatternExtractor } from "../extracter/patterns/date-pattern-extractor";
import { createAInvestmentLog } from "./stock-slip/core";
import { TransactionExtractor } from "./transaction/transaction-extractor";

export function dimeManagerParser(text: string) {
    if (text.includes('Stock Amount')) return { slip: [createAInvestmentLog(text).toJson()] }
    const dateExtractor = new DatePatternExtractor();
    const extractor = new TransactionExtractor(dateExtractor, text);
    return { ...extractor.toJson() }
} 