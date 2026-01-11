import express, { type Express, type Request, type Response } from "express";
import cors from "cors";
import { TaskManager } from "../services/task/task";
import v1DimeRoute from "./v1-routes/dime";
import v2BinanceThRoute from "./v2-routes/binance-th";
import v2DimeRoute from "./v2-routes/dime";
import v1BinanceThRoute from "./v1-routes/binance-th";
let app: Express = express();
const port: number = 8080;
app.use(cors());
app.use(express.json());

app.use((req, res, next) => {
  const requestTime = new Date(Date.now()).toISOString();
  console.log(`[${requestTime}] ${req.method} ${req.url}`);
  next();
});
app.get("/", (req: Request, res: Response) => {
  res.send("Hello, Express with TypeScript!");
});
export function profilerMiddleware(req: any, res: any, next: Function) {
  const startTime = process.hrtime.bigint();
  const startUsage = process.cpuUsage();
  const startMemory = process.memoryUsage().heapUsed;

  res.on('finish', () => {
    const endTime = process.hrtime.bigint();
    const endUsage = process.cpuUsage(startUsage);
    const endMemory = process.memoryUsage().heapUsed;

    const elapsedTimeMs = Number(endTime - startTime) / 1e6;
    const cpuTotalMs = (endUsage.user + endUsage.system) / 1000;
    const cpuPercent = ((cpuTotalMs / elapsedTimeMs) * 100).toFixed(2);
    const millicores = ((cpuTotalMs / elapsedTimeMs) * 1000).toFixed(0);
    const memoryDiff = ((endMemory - startMemory) / 1024 / 1024).toFixed(3);

    console.log(`--- Profiler: ${req.method} ${req.originalUrl} ---`);
    console.log(`⏱️  Response Time: ${elapsedTimeMs.toFixed(2)} ms`);
    console.log(`💻 CPU Power Used: ${millicores}m (จาก 1000m ต่อ 1 Core)`);
    console.log(`📊 CPU Efficiency: ${cpuPercent}%`);
    console.log(`🧠 RAM Used: ${memoryDiff} MB`);
    console.log('-----------------------------------\n');
  });

  next();
};
export function loggerMiddleWare(err: any, req: Request, res: Response, next: Function) {
  console.error(err.message)
  const statusCode = err.statusCode || 500;
  res.status(statusCode).send({
    status: statusCode,
    message: err.message || 'Internal Server Error',
  });
}
const tasks = new TaskManager(20)
app.use(profilerMiddleware)
app.use(loggerMiddleWare);
app = v1BinanceThRoute(app)
app = v2DimeRoute(app, tasks)
app = v1DimeRoute(app)
app = v2BinanceThRoute(app, tasks)
app.use(express.json());
try {
  app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
  });
}
catch (ex) {
  console.error(ex)
}
