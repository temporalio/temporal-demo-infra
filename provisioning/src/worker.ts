import fs from 'fs/promises';

// @@@SNIPSTART typescript-mtls-worker
import { Worker, NativeConnection } from '@temporalio/worker';
import * as activities from './activities';
import 'dotenv/config'

/**
 * Run a Worker with an mTLS connection, configuration is provided via environment variables.
 * Note that serverNameOverride and serverRootCACertificate are optional.
 */
async function run({
                     address,
                     namespace,
                     clientCertPath,
                     clientKeyPath,
                     serverNameOverride,
                     serverRootCACertificatePath,
                     taskQueue,
                   }: Env) {
  let serverRootCACertificate: Buffer | undefined = undefined;
  if (serverRootCACertificatePath) {
    serverRootCACertificate = await fs.readFile(serverRootCACertificatePath);
  }

  const connection = await NativeConnection.connect({
    address,
    tls: {
      serverNameOverride,
      serverRootCACertificate,
      // See docs for other TLS options
      clientCertPair: {
        crt: await fs.readFile(clientCertPath),
        key: await fs.readFile(clientKeyPath),
      },
    },
  });

  const worker = await Worker.create({
    connection,
    namespace,
    // workflowsPath: require.resolve('./workflows'),
    activities,
    taskQueue,
  });
  console.log('Worker connection successfully established');

  await worker.run();
  await connection.close();
}

run(getEnv()).catch((err) => {
  console.error(err);
  process.exit(1);
});
// @@@SNIPEND

// Helpers for configuring the mTLS client and worker samples

function requiredEnv(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new ReferenceError(`${name} environment variable is not defined`);
  }
  console.log('found',name, value)
  return value;
}

export interface Env {
  address: string;
  namespace: string;
  clientCertPath: string;
  clientKeyPath: string;
  serverNameOverride?: string;
  serverRootCACertificatePath?: string;
  taskQueue: string;
}

export function getEnv(): Env {
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