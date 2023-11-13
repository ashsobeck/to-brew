<script lang="ts">
	let showBrew = false;
</script>

{#if showBrew}
	<form method="POST" action="?/create">
		<div class="card p-2">
			<label class="label"><input class="input" name=""/></label>
			<label class="label">Weight: <input class="input" name="weight" /> </label>
			<label class="label">Time of Brew: <input class="input" name="time" /> </label>
			<button class="button" type="submit">Create</button>
		</div>
	</form>
{/if}
<button disabled={showBrew} on:click={() => (showBrew = true)}>+ Add Brew</button>
