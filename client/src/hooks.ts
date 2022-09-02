import process from "process";

process.on('SIGINT', function () {process.exit();}); // Ctrl+C  
process.on('SIGTERM', function () {process.exit();}); // docker stop
