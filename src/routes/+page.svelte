<script lang="ts">
	import { writable, type Writable } from 'svelte/store';
	import type { ToBrew } from '../lib/types';
	import ToBrewList from './ToBrewList.svelte';

	const store = makeStore();
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<ToBrewList {store} done={false} />
