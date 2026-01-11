import type { Parser } from "../parser";
import { BuyStockTransaction } from "./buy-stock";
import { SellStockTransaction } from "./sell-stock";
import type { Stock } from "./stock";

export function getStockParser(text: string): Parser<Stock> {
  if (text.startsWith("Sell")) return new SellStockTransaction(text);
  if (text.startsWith("Buy")) return new BuyStockTransaction(text);
  throw new Error("no stock parser found");
}
