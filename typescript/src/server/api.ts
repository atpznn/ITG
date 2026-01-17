import express, { type Express, type Request, type Response } from "express";
import cors from "cors";
import { TaskManager } from "../services/task/task";
import { profilerIntercepter } from "./intercepters/profiler.intercepter";
import { requestMiddleWare } from "./middlewares/request.middleware";
import { dimeManagerParser } from "../services/dime/manager";
import imageProcessRoute, { doOcrImage, doTask } from "./routes/image-process";
import { cleanText } from "../services/ocr";
import { binanceThManagerParser } from "../services/binance-th/manager";
import { errorIntercepter } from "./intercepters/error.intercepter";
import taskRoute from "./routes/task";
let app: Express = express();
const port: number = 8080;
app.use(cors());
app.use(express.json());
app.get("/", (req: Request, res: Response) => res.send("Hello, Express with TypeScript!"));
// const tasks = new TaskManager(20)
// app.use(profilerIntercepter)
// app.use(requestMiddleWare);
// app.use(errorIntercepter);
// app.use('/v1/dime', imageProcessRoute(doOcrImage(dimeManagerParser)))
// app.use('/v1/dime/process-text', (req: Request, res: Response) => res.status(200).json((dimeManagerParser(cleanText(req.body.text)))))
// app.use('/v1/binance-th', imageProcessRoute(doOcrImage(binanceThManagerParser)))
// app.use('/v2/binance-th', (req: Request) => imageProcessRoute(doTask(tasks, req.baseUrl.replace('/', ''), binanceThManagerParser)))
// app.use('/v2/dime', (req: Request) => imageProcessRoute(doTask(tasks, req.baseUrl.replace('/', ''), dimeManagerParser)))
// app.use(taskRoute(tasks))
app.use(express.json());
app.use(imageProcessRoute(doOcrImage(text => text)))
app.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});
