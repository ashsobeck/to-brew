<script lang="ts">
	import type { ToBrew } from '$lib/types';
	import { tobrews } from '$lib/tobrews';
	import AddBrew from './AddBrew.svelte';

	let brews: ToBrew[];

	tobrews.subscribe((b) => {
		console.log(b);
		brews = b;
	});

	$: toBrew = brews.filter((b: ToBrew) => b.brewed === false);
	$: brewed = brews.filter((b: ToBrew) => b.brewed === true);
</script>

<div class="flex p-2 w-full space-x-2">
	<div class="card p-4 basis-1/2">
		To Brew
		<ul>
			{#each toBrew as brew, i}
				<div class="p-2">
					<li>{brew.bean}</li>
					<li>{brew.name}</li>
					<li>{brew.id}</li>
					<li>{new Date(brew.time).toTimeString()}</li>
					<li>{brew.brewed}</li>
					<li>{i}</li>
				</div>
			{/each}
			<AddBrew />
		</ul>
	</div>
	<div class="card p-4 basis-1/2">
		Brewed
		<ul>
			{#each brewed as brew, i}
				<div class="p-2">
					<li>{brew.bean}</li>
					<li>{brew.name}</li>
					<li>{brew.id}</li>
					<li>{brew.time}</li>
					<li>{brew.brewed}</li>
					<li>{i}</li>
				</div>
			{/each}
		</ul>
	</div>
</div>
