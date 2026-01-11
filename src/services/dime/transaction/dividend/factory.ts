import { IncomeDividendLog } from "./income-dividend";
import { TaxDividendLog } from "./tax-dividend";

export function getDividendParser(text: string) {
  const type = text!.includes("Withholding Tax") ? "Tax" : "Income";
  switch (type) {
    case "Income": {
      return new IncomeDividendLog(text);
    }
    case "Tax": {
      return new TaxDividendLog(text);
    }
    default:
      throw Error("no type");
  }
}
