<script lang="ts">
	import { beans, tobrews } from '$lib/stores';
	import type { Bean, ToBrew } from '$lib/types';
	import type { ActionData } from './$types';
	import ToBrewList from './ToBrewList.svelte';
	type BrewData = {
		brews: ToBrew[];
		beans: Bean[];
	};
	export let data: BrewData;
	export let form: ActionData;

	console.log('getting brews...');
	tobrews.set(data.brews);
	beans.set(data.beans);
</script>

<ToBrewList {form} />
