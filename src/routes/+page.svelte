<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew } from '$lib/types';
	import ToBrewList from './ToBrewList.svelte';
	type BrewData = {
		brews: ToBrew[];
	};
	export let data: BrewData;

	console.log(data);
	tobrews.set(data.brews);
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<ToBrewList />
