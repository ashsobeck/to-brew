<script lang="ts">
	import type { ToBrew } from '$lib/types';
	import { tobrews } from '$lib/tobrews';
	import AddBrew from './AddBrew.svelte';

	let brews: ToBrew[];

	tobrews.subscribe((b) => {
		console.log(b);
		brews = b;
	});

	$: toBrew = brews.filter((b: ToBrew) => b.done === true);
	$: brewed = brews.filter((b: ToBrew) => b.done === false);
</script>

<div class="card p-4">
	<ul>
		{#each toBrew as brew, i}
			<div class="card">
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
</div>
<div class="card p-4">
	<ul>
		{#each brewed as brew, i}
			<div class="card">
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
</div>
