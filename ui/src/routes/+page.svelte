<script lang="ts">
    import '../app.css'
    import Button  from "$lib/holocene/button.svelte";
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
    let profile: string
    let result: string
    const _submitApp = async () => {
        const res = await fetch('http://localhost:8042/api/provision', {
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
        const res = await fetch('http://localhost:8042/api/provision', {
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
        } else {
            console.error(res.status)
        }

    }
    const _rejectApp = async () => {
        console.log('authorizing', submitted)
        const res = await fetch('http://localhost:8042/api/provision', {
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
        } else {
            console.error(res.status)
        }

    }

</script>

<div class="min-h-screen bg-primary text-offWhite flex flex-col text-xl items-center justify-center">
    <header class="text-offWhite justify-center align-center flex flex-col">
        <h1 class="text-center text-4xl">Provision Application</h1>
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
        <h3>Approve "{ submitted.applicationName }"?</h3>
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
    {/if}
</div>

<style lang="postcss">
    button {
        background: green;
    }
</style>