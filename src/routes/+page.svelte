<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew } from '$lib/types';
	import type { ActionData } from './$types';
	import ToBrewList from './ToBrewList.svelte';
	type BrewData = {
		brews: ToBrew[];
	};
	export let data: BrewData;
	export let form: ActionData;

	console.log('getting brews...');
	tobrews.set(data.brews);
</script>

<ToBrewList {form} />
