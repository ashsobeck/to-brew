<script lang="ts">
	import type { ToBrew } from '$lib/types';
	import { tobrews } from '$lib/tobrews';

	let brews: ToBrew[];

	tobrews.subscribe((b) => {
		console.log(b);
		brews = b;
	});
</script>

<ul>
	{#each brews as brew, i}
		<div>
			<div>{brew.bean}</div>
			<div>{brew.name}</div>
			<div>{brew.id}</div>
			<div>{brew.time}</div>
			<div>{brew.done}</div>
			<div>{i}</div>
		</div>
	{/each}
</ul>
