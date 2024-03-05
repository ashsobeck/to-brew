<script lang="ts">
	import type { ToBrew } from '$lib/types';
	import { tobrews } from '$lib/stores';
	import AddBrew from './AddBrew.svelte';
	import { applyAction, enhance } from '$app/forms';
	import type { ActionData } from './$types';

	export let form: ActionData;

	$: toBrew = $tobrews.filter((b: ToBrew) => b.brewed === false);
	$: brewed = $tobrews.filter((b: ToBrew) => b.brewed === true);
	console.log(toBrew);
	console.log(brewed);

	let updateBrew = () => {
		return async ({ result }) => {
			console.log(result);
			console.log($tobrews);
			// needed to update the form prop
			await applyAction(result);
			const brewIdx = $tobrews.findIndex((brew) => brew.id === form?.brew.id);
			console.log('form:');
			console.log(form);
			console.log(brewIdx);

			$tobrews[brewIdx] = form?.brew as ToBrew;

			$tobrews = [...$tobrews];
			console.log($tobrews);
		};
	};
</script>

<div class="flex p-2 w-full space-x-2">
	<div class="card p-4 basis-1/2">
		To Brew
		<ul>
			{#each toBrew as brew, i}
				<div class="card card-hover p-2">
					<li>{brew.bean}</li>
					<!-- <li>{brew.name}</li> -->
					<li>{brew.id}</li>
					<li>{brew.timeToBrew}</li>
					<li>{brew.brewed}</li>
					<li>{i}</li>
					<form method="POST" action="?/brewed" use:enhance={updateBrew}>
						<input type="hidden" name="id" value={brew.id} />
						<button class="btn variant-filled">brew</button>
					</form>
				</div>
			{/each}
			<AddBrew />
		</ul>
	</div>
	<div class="card card-hover p-4 basis-1/2">
		Brewed
		<ul>
			{#each brewed as brew, i}
				<div class="card card-hover p-2">
					<li>{brew.bean}</li>
					<!-- <li>{brew.name}</li> -->
					<li>{brew.id}</li>
					<li>{brew.timeToBrew}</li>
					<li>{brew.brewed}</li>
					<li>{i}</li>
				</div>
				<form method="POST" action="?/update" use:enhance={updateBrew}>
					<input type="hidden" name="id" value={brew.id} />
					<!-- <input type="hidden" name="name" value={brew.name} /> -->
					<input type="hidden" name="time" value={new Date(brew.timeToBrew).toISOString()} />
					<!-- <input type="hidden" name="link" value={brew.link} /> -->
					<!-- <input type="hidden" name="roaster" value={brew.roaster} /> -->
					<input type="hidden" name="bean" value={brew.bean} />
					<input type="hidden" name="brewed" value={brew.brewed} />
					<button class="btn variant-filled">rebrew</button>
				</form>
			{/each}
		</ul>
	</div>
</div>
