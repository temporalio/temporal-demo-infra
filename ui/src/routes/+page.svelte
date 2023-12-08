<script lang="ts">
    import '../app.css'
    import Button  from "$lib/holocene/button.svelte";
    import { PUBLIC_API_ROOT_HOST, PUBLIC_API_ROOT_SCHEME } from '$env/static/public'

    interface ProvisionResponse {
        workflowId: string
        authorizerId: string
        teamId: string
        applicationName: string
        applicationId: string
    }
    let teamId: string;
    let applicationName: string;
    let message: string;
    let authorizerId: string
    let region: string
    let submitted: ProvisionResponse | undefined = undefined
    let approved: boolean
    let rejected: boolean
    let profile: string
    let result: string
    let provisionURL: string = `${PUBLIC_API_ROOT_SCHEME}${PUBLIC_API_ROOT_HOST}/provision`
    const temporalLogoUrl = new URL('../static/temporal_logo.png', import.meta.url).href

    const _submitApp = async () => {
        const res = await fetch(provisionURL, {
            method: 'POST',
            body: JSON.stringify({
                applicationName,
                teamId,
                authorizerId,
            }),
            headers: {
                'content-type': 'application/json'
            }
        })

        if(res.ok) {
            console.log('success',res.status)
            const json = await res.json()
            submitted = json
            authorizerId = submitted?.authorizerId
            applicationName = submitted?.applicationName
        } else {
            console.error(res.status)
        }
    }
    const _authorizeApp = async () => {
        console.log('authorizing', submitted)
        const res = await fetch(provisionURL, {
            method: 'PATCH',
            body: JSON.stringify({
                region: region,
                profile: profile,
                workflowId: submitted?.workflowId,
                applicationId: submitted?.applicationId
            }),
            headers: {
                'content-type': 'application/json'
            }
        })

        if(res.ok) {
            console.log('success',res.status)
            approved = true
        } else {
            console.error(res.status)
        }

    }
    const _rejectApp = async () => {
        console.log('authorizing', submitted)
        const res = await fetch(provisionURL, {
            method: 'DELETE',
            body: JSON.stringify({
                region: region,
                profile: profile,
                workflowId: submitted?.workflowId,
                applicationId: submitted?.applicationId,
            }),
            headers: {
                'content-type': 'application/json'
            }
        })

        if(res.ok) {
            console.log('success',res.status)
            rejected = true
        } else {
            console.error(res.status)
        }

    }

</script>

<div class="min-h-screen bg-primary text-offWhite flex flex-col text-xl items-center justify-center">

    <header class="text-offWhite justify-center align-center flex flex-col">
        <h1 class="text-center flex justify-center items-center">
            <img class="h-auto w-auto max-x-full max-w-12 w-12 rounded-full" src="{temporalLogoUrl}" alt="Temporal"/>
        </h1>
        <h1 class="text-center text-4xl">Temporal Platform</h1>
        <hr class="h-px my-8 bg-gray-200 border-0 dark:bg-gray-700">
        <h2 class="text-center text-2xl">Provision Application</h2>
        <hr class="h-px my-8 bg-gray-200 border-0 dark:bg-gray-700">
    </header>
    {#if !submitted}
        <div class="flex flex-col text-inherit">
            <label class="text-inherit flex flex-col m-4 w-96" for="app_name">
                <span>Application Name</span>
                <input class="text-primary p-2" name="app_name" type="text" bind:value={ applicationName } />
            </label>
            <label class="text-inherit flex flex-col m-4 w-96" for="team_id">
                <span>Team Id</span>
                <input class="text-primary p-2" name="team_id" type="text" bind:value={ teamId } />
            </label>
            <label class="text-inherit flex flex-col m-4 w-96" for="authorizer_id">
                <span>Authorizer Id</span>
                <input class="text-primary p-2" name="authorizer_id" type="text" bind:value={ authorizerId } />
            </label>


            <Button class="self-center border-spaceGray" variant="secondary" on:click={ _submitApp }>Submit Provision Request</Button>
            { #if message }
                <p class="error message">{ message }</p>
            { /if }

        </div>
    {/if}
    { #if submitted}
        { #if approved }
            <div class="p-4 mb-4 text-sm text-green-800 rounded-lg bg-green-50 dark:bg-gray-800 dark:text-green-400" role="alert">
                <span class="font-medium">APPROVED!</span> The application is being provisioned. Check the Temporal UI for progress...
            </div>
        {:else if rejected }
            <div class="p-4 mb-4 text-sm text-yellow-800 rounded-lg bg-yellow-50 dark:bg-gray-800 dark:text-yellow-300" role="alert">
                <span class="font-medium">REJECTED</span> Resubmit the application for authorization later.
            </div>
        { :else}
            <h3 class="p-4 mb-4 text-sm text-blue-800 rounded-lg bg-blue-50 dark:bg-gray-800 dark:text-blue-400" role="alert">
                <span class="font-medium">APPROVAL NEEDED:</span> Should the application "{ submitted.applicationName }" be provisioned for team { submitted.teamId }?
            </h3>

            <div class="flex flex-col text-inherit">
                <label class="text-inherit flex flex-col m-4 w-96" for="region">
                    <span>Region</span>
                    <input class="text-primary p-2" name="region" type="text" bind:value={ region } />
                </label>
                <label class="text-inherit flex flex-col m-4 w-96" for="profile">
                    <span>Profile</span>
                    <input class="text-primary p-2" name="profile" type="text" bind:value={ profile } />
                </label>
                <label class="text-inherit flex flex-col m-4 w-96" for="authorizer_id">
                    <span>Authorizer ID</span>
                    <input class="text-primary p-2" name="authorizer_id" type="text" bind:value={ authorizerId } />
                </label>


                <Button class="self-center border-spaceGray" variant="secondary" on:click={ _authorizeApp }>APPROVE</Button>
                <Button class="self-center border-spaceGray" variant="secondary" on:click={ _rejectApp }>REJECT</Button>
                { #if message }
                    <p class="error message">{ message }</p>
                { /if }
            </div>
        { /if }

    {/if}
</div>

<style lang="postcss">
    button {
        background: green;
    }

</style>