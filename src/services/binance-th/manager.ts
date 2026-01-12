import { BinanceThTransactionPatternExtractor } from "../extracter/patterns/binance-th-transaction-pattern-extractor";
import { BinanceThSlip } from "./slip/slip";
import { BinanceThTransaction } from "./transaction/transaction";
export function binanceThManagerParser(text: string) {
    const dateExtractor = new BinanceThTransactionPatternExtractor();
    if (text.toLowerCase().includes('details')) {
        return { slip: [new BinanceThSlip(text).toJson()] }
    }
    const extractor = new BinanceThTransaction(dateExtractor, text);
    return { transaction: extractor.toJson() }
}