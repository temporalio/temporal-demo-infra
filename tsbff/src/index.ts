import express, { Express, Request } from "express";
import 'dotenv/config'
import { getClient} from "./client";
import cuid from "cuid";


// Application Setup
const app: Express = express();
const port = 3000;
interface ProvisionApplicationRequest{
    requesterId: string
    applicationId: string
    teamId: string
    applicationName: string
    authorizerId: string
    authorizationTimeoutSeconds: number
    demoAuthorizationDelaySeconds: number
}

app.post('/provision',  (req:Request<ProvisionApplicationRequest,ProvisionApplicationRequest>, res) => {
    console.log('received req', req.body)
    getClient().then(tc => {
        console.log('got client')
        let args: ProvisionApplicationRequest = {
            applicationId: cuid(),
            applicationName: req.body.applicationName,
            authorizationTimeoutSeconds: 120,
            authorizerId: cuid(),
            demoAuthorizationDelaySeconds: 0,
            teamId: req.body.teamId,
            requesterId: cuid()
        }
        console.log('starting it', req.params)
        return tc.workflow.start('ProvisionApplication', {
            args: [args], taskQueue: "apps", workflowId: "monsters"
        }).then((run) => {
            console.log('runId', run.firstExecutionRunId)
            res.status(200)
        }).catch(err => {
            console.error(err)
        })
    }).catch(err => {
        console.error('failed to get Client', err)
    })
})
app.listen(port, () => {
    console.log(`[server]: Server is running at http://localhost:${port}`);
});