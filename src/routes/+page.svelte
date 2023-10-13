<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew } from '$lib/types';
	import ToBrewList from './ToBrewList.svelte';
	type BrewData = {
		brews: ToBrew[];
	};
	export let data: BrewData;

	console.log('getting brews...');
	tobrews.set(data.brews);
</script>

<ToBrewList />
