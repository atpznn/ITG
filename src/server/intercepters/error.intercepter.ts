import type { Request, Response } from "express";

export function errorIntercepter(err: any, req: Request, res: Response, next: Function) {
    console.error(err.message)
    const statusCode = err.statusCode || 500;
    res.status(statusCode).send({
        status: statusCode,
        message: err.message || 'Internal Server Error',
    });
}