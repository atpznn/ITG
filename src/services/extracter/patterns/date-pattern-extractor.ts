import { BasePatternExtractor } from "./base-pattern-extractor";

export class DatePatternExtractor extends BasePatternExtractor {
  constructor() {
    super(/(?<=\d{1,2}\s[A-Z][a-z]{2}\s\d{4}\s-\s\d{2}:\d{2}:\d{2}\s[AP]M)/g);
  }
}
