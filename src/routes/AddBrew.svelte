<script lang="ts">
	import { beans } from '$lib/stores';
	import { SlideToggle } from '@skeletonlabs/skeleton';

	let showBrew = false;
	let defaultTime = new Date().toLocaleTimeString([], {
		hour12: false,
		hour: '2-digit',
		minute: '2-digit'
	});

	$: beanNamesAndIds = $beans.map((b) => {
		return { name: b.name, id: b.id };
	});
</script>

{#if showBrew}
	<form method="POST" action="?/create">
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
