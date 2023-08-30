<script lang="ts">
	import type { ToBrew } from '$lib/types';
	import { tobrews } from '$lib/tobrews';
	import AddBrew from './AddBrew.svelte';

	let brews: ToBrew[];

	tobrews.subscribe((b) => {
		console.log(b);
		brews = b;
	});
</script>

<ul>
	{#each brews as brew, i}
		<div>
			<li>{brew.bean}</li>
			<li>{brew.name}</li>
			<li>{brew.id}</li>
			<li>{brew.time}</li>
			<li>{brew.done}</li>
			<li>{i}</li>
		</div>
	{/each}
	<AddBrew />
</ul>
