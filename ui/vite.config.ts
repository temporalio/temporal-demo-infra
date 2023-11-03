import { sveltekit } from '@sveltejs/kit/vite'
import type { UserConfig } from 'vite'
import path from 'path';
import mkcert from 'vite-plugin-mkcert'


/** @type {import('vite').UserConfig} */
const config: UserConfig = {
	base: '/viteconfignotused',
	plugins: [mkcert(),  sveltekit()],
	resolve: {
		alias: {
			$lib: path.resolve(__dirname, './src/lib'),
			$log: path.resolve(__dirname, './src/lib/log'),
			$gql: path.resolve(__dirname, './src/gql'),
			$types: path.resolve(__dirname, './src/types'),
			$components: path.resolve(__dirname, './src/lib/components/'),
			//$app: path.resolve(__dirname, './src/lib/svelte-mocks/app/'),
			$fixtures: path.resolve(__dirname, './src/fixtures/'),
			$holocene: path.resolve(__dirname, './src/lib/holocene'),
		},
	},
	build: {
		sourcemap: true
	},
	optimizeDeps: {
		exclude: ['@urql/svelte']
	}
}

export default config
