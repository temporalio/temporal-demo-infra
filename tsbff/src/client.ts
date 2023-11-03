import fs from 'fs/promises';

// @@@SNIPSTART typescript-mtls-worker
import { Connection, Client } from '@temporalio/client';
import 'dotenv/config'




async function createClient({
                     address,
                     namespace,
                     clientCertPath,
                     clientKeyPath,
                     serverNameOverride,
                     serverRootCACertificatePath,
                     taskQueue,
                   }: Env) {
  // Note that the serverRootCACertificate is NOT needed if connecting to Temporal Cloud because
  // the server certificate is issued by a publicly trusted CA.
  let serverRootCACertificate: Buffer | undefined = undefined;
  if (serverRootCACertificatePath) {
    serverRootCACertificate = await fs.readFile(serverRootCACertificatePath);
  }

  const connection = await Connection.connect({
    address,
    tls: {
      serverNameOverride,
      serverRootCACertificate,
      clientCertPair: {
        crt: await fs.readFile(clientCertPath),
        key: await fs.readFile(clientKeyPath),
      },
    },
  });
  const client = new Client({ connection, namespace });
  return client
}
let defaultClient: Client
export const getClient = async () => {
    if(defaultClient) {
        return defaultClient
    }
    defaultClient = await createClient(getEnv()).then(
        () => process.exit(0),
        (err) => {
            console.error(err);
            process.exit(1);
        })
    return defaultClient
}


// @@@SNIPEND

// Helpers for configuring the mTLS client and worker samples

function requiredEnv(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new ReferenceError(`${name} environment variable is not defined`);
  }
  return value;
}

export interface Env {
  address: string;
  namespace: string;
  clientCertPath: string;
  clientKeyPath: string;
  serverNameOverride?: string; // not needed if connecting to Temporal Cloud
  serverRootCACertificatePath?: string; // not needed if connecting to Temporal Cloud
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