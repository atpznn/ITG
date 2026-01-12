import multer from "multer";
import express, {
  type Request,
  type Response,
} from "express";
import { parseImageToText } from "../../services/ocr";
import { CoordinatesOcrStategy } from "../../services/ocr/stategies/coordinates-ocr";
import type { TaskManager } from "../../services/task/task";
const upload = multer({
  limits: { fileSize: 5 * 1024 * 1024 },
}).array("images", 5);

function imageHandler(doThing: (f: Express.Multer.File[]) => Promise<unknown>) {
  return function (req: Request, res: Response) {
    return upload(req, res, async (err) => {
      if (err instanceof multer.MulterError) return res.status(400).send(`Multer Error: ${err.message}`);
      if (err) return res.status(500).send("An unknown error occurred during upload.");
      const files = (req as any).files as Express.Multer.File[]
      if (!files || files.length === 0) return res.status(400).send("No files were uploaded.");

      console.log('found image file :', files.length)
      const result = await doThing(files)
      res.json(result)
    })
  }
}
export default function imageProcessRoute(doThing: (f: Express.Multer.File[]) => Promise<unknown>) {
  const app = express();
  app.post('/image-process', imageHandler(doThing));
  return app
}
export function doOcrImage(doAfterOcr: (text: string) => unknown) {
  return async function (files: Express.Multer.File[]) {
    const results = []
    for (const file of files) {
      const text = await parseImageToText(file.buffer, new CoordinatesOcrStategy())
      results.push(doAfterOcr(text))
    }
    return results
  }
}
export function doTask(taskManager: TaskManager, taskName: string, doAfterOcr: (text: string) => unknown) {
  return async function (files: Express.Multer.File[]) {
    const todos: (() => Promise<unknown>)[] = []
    for (const file of files) {
      todos.push(async () => {
        const text = await parseImageToText(file.buffer, new CoordinatesOcrStategy())
        return doAfterOcr(text)
      })
    }
    const taskId = taskManager.spawnNewTask(taskName, todos)
    return taskId
  }
}
