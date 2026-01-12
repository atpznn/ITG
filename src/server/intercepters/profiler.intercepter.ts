export function profilerIntercepter(req: any, res: any, next: Function) {
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