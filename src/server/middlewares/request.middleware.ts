import type { Request, Response } from "express";

export function requestMiddleWare(err: any, req: Request, res: Response, next: Function) {
    const requestTime = new Date(Date.now()).toISOString();
    console.log(`[${requestTime}] ${req.method} ${req.url}`);
    next();
}