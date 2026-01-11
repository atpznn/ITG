import type { Express } from "express";
import { BinanceThTransactionPatternExtractor } from "../../services/extracter/patterns/binance-th-transaction-pattern-extractor";
import { BinanceThSlip } from "../../services/binance-th/slip/slip";
import { BinanceThTransaction } from "../../services/binance-th/transaction/transaction";
import imageProcessRoute from "./image-process";
import { TaskManager } from "../../services/task/task";
const binanceThHandler = (text: string) => {
    try {
        const dateExtractor = new BinanceThTransactionPatternExtractor();
        if (text.toLowerCase().includes('details')) {
            return { slip: [new BinanceThSlip(text).toJson()] }
        }
        const extractor = new BinanceThTransaction(dateExtractor, text);
        return { transaction: extractor.toJson() }
    }
    catch (ex) {
        throw ex
    }
};
export default function v2BinanceThRoute(app: Express, taskManager: TaskManager) {
    app.use("/v2/binance-th", imageProcessRoute(taskManager, binanceThHandler));
    return app
}
