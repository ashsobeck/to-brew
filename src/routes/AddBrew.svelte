<script lang="ts">
	let showBrew = false;
</script>

{#if showBrew}
	<form method="POST" action="?/create">
		<div class="card p-2">
			<label class="label"
				>Bean:
				<select class="select" name="bean">
					<option value="big bean"> bean bean </option>
				</select>
			</label>
			<label class="label">Weight: <input class="input" name="weight" /> </label>
			<label class="label">Time of Brew: <input class="input" name="time" /> </label>
			<!-- <label class="label">Name: <input class="input" name="name" /> </label> -->
			<!-- <label class="label">Roaster: <input class="input" name="roaster" /> </label> -->
			<!-- <label class="label">Link: <input class="input" name="link" /> </label> -->
			<button class="button" type="submit">Create</button>
		</div>
	</form>
{/if}
<button disabled={showBrew} on:click={() => (showBrew = true)}>+ Add Brew</button>
