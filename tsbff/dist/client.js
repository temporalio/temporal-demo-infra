var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
import fs from 'fs/promises';
// @@@SNIPSTART typescript-mtls-worker
import { Connection, Client } from '@temporalio/client';
import 'dotenv/config';
function createClient(_a) {
    var address = _a.address, namespace = _a.namespace, clientCertPath = _a.clientCertPath, clientKeyPath = _a.clientKeyPath, serverNameOverride = _a.serverNameOverride, serverRootCACertificatePath = _a.serverRootCACertificatePath, taskQueue = _a.taskQueue;
    return __awaiter(this, void 0, void 0, function () {
        var serverRootCACertificate, connection, _b, _c, client;
        var _d, _e, _f;
        return __generator(this, function (_g) {
            switch (_g.label) {
                case 0:
                    serverRootCACertificate = undefined;
                    if (!serverRootCACertificatePath) return [3 /*break*/, 2];
                    return [4 /*yield*/, fs.readFile(serverRootCACertificatePath)];
                case 1:
                    serverRootCACertificate = _g.sent();
                    _g.label = 2;
                case 2:
                    _c = (_b = Connection).connect;
                    _d = {
                        address: address
                    };
                    _e = {
                        serverNameOverride: serverNameOverride,
                        serverRootCACertificate: serverRootCACertificate
                    };
                    _f = {};
                    return [4 /*yield*/, fs.readFile(clientCertPath)];
                case 3:
                    _f.crt = _g.sent();
                    return [4 /*yield*/, fs.readFile(clientKeyPath)];
                case 4: return [4 /*yield*/, _c.apply(_b, [(_d.tls = (_e.clientCertPair = (_f.key = _g.sent(),
                            _f),
                            _e),
                            _d)])];
                case 5:
                    connection = _g.sent();
                    client = new Client({ connection: connection, namespace: namespace });
                    return [2 /*return*/, client];
            }
        });
    });
}
var defaultClient;
export var getClient = function () { return __awaiter(void 0, void 0, void 0, function () {
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                if (defaultClient) {
                    return [2 /*return*/, defaultClient];
                }
                return [4 /*yield*/, createClient(getEnv()).then(function () { return process.exit(0); }, function (err) {
                        console.error(err);
                        process.exit(1);
                    })];
            case 1:
                defaultClient = _a.sent();
                return [2 /*return*/, defaultClient];
        }
    });
}); };
// @@@SNIPEND
// Helpers for configuring the mTLS client and worker samples
function requiredEnv(name) {
    var value = process.env[name];
    if (!value) {
        throw new ReferenceError("".concat(name, " environment variable is not defined"));
    }
    return value;
}
export function getEnv() {
    return {
        address: requiredEnv('TEMPORAL_ADDRESS'),
        namespace: requiredEnv('TEMPORAL_NAMESPACE'),
        clientCertPath: requiredEnv('TEMPORAL_CLIENT_CERT_PATH'),
        clientKeyPath: requiredEnv('TEMPORAL_CLIENT_KEY_PATH'),
        serverNameOverride: process.env.TEMPORAL_SERVER_NAME_OVERRIDE,
        serverRootCACertificatePath: process.env.TEMPORAL_SERVER_ROOT_CA_CERT_PATH,
        taskQueue: process.env.TEMPORAL_TASK_QUEUE || 'provisioning_aws',
    };
}
