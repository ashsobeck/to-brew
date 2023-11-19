<script lang="ts">
	import { beans } from '$lib/stores';

	let showBrew = false;

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
			<label class="label">Time of Brew: <input class="input" name="time" /> </label>
			<button class="button" type="submit">Create</button>
		</div>
	</form>
{/if}
<button disabled={showBrew} on:click={() => (showBrew = true)}>+ Add Brew</button>
