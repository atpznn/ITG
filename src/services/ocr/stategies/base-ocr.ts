import tesseract from "node-tesseract-ocr"
export abstract class OcrStategy {
    abstract getConfig: () => tesseract.Config
    abstract mutation: (str: string) => string
}
export class BaseOcrStategy implements OcrStategy {
    private config: tesseract.Config = {
        lang: "eng+tha",
        oem: 1,
        psm: 6,

        // "textord_heavy_nr": "0",
        // "textord_min_linesize": '1.0',
        // "tessedit_do_invest": "0",
        // "paragraph_text_based": "0",
        // "tessdata_manager_debug_level": "0"

    }
    public getConfig(): tesseract.Config {
        return this.config
    }
    public mutation(str: string) {
        return str
    }
}