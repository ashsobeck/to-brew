<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import { beans, tobrews } from '$lib/stores';
	import type { ToBrew } from '$lib/types';
	import type { ActionData } from './$types';

	export let form: ActionData;
	let showBrew = false;
	let defaultTime = new Date().toLocaleTimeString([], {
		hour12: false,
		hour: '2-digit',
		minute: '2-digit'
	});

	$: beanNamesAndIds = $beans.map((b) => {
		return { name: b.name, id: b.id };
	});

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

{#if showBrew}
	<form method="POST" action="?/create" use:enhance={updateBrew}>
		<div class="card p-2">
			<label class="label"
				>Bean:
				<select class="select" name="bean">
					{#each beanNamesAndIds as nameId}
						<option value={nameId.id}>{nameId.name}</option>
					{/each}
				</select>
			</label>
			<label class="label">Weight: <input class="input" name="weight" /> </label>
			<label class="label"
				>Time of Brew:
				<div class="flex flex-row">
					<input class="input" name="time" type="time" value={defaultTime} />
					<!-- <input class="input" type="number" maxlength="2" /> -->
					<!-- <SlideToggle name="slider-label" on:change={() => (pastMidday = !pastMidday)}> -->
					<!-- 	{pastMidday ? 'PM' : 'AM'} -->
					<!-- </SlideToggle> -->
				</div>
			</label>
			<button class="button" type="submit">Create</button>
		</div>
	</form>
{/if}
<button disabled={showBrew} on:click={() => (showBrew = true)}>+ Add Brew</button>
