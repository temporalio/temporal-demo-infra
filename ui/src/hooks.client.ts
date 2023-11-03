// import type {HandleServerError} from "@sveltejs/kit";
// import {redirect} from "@sveltejs/kit";
//
export const handleError= ({ error, event }) => {
    console.log('hooks.client.ts','#handleError', 'error', error, 'event', event)
    // if(error.code == 401) {
    //     console.log('err', error)
    //     event.cookies.set("session", "", { path: "/", httpOnly: true, secure: true, maxAge: 0 });
    //     // return redirect(302,'/login')
    // }
}
