<script lang="ts">
	import { beans } from '$lib/stores';
</script>

<div class="card">
	{#each $beans as bean}
		<div class="card">
			<h3>{bean.name}</h3>
			<ul>
				<li>{bean.roaster}</li>
				<li>{bean.country}</li>
				<li>{bean.varietal}</li>
				<li>{bean.process}</li>
				<li>{bean.altitude}</li>
				<li>Amount left {bean.weight}g</li>
			</ul>
		</div>
	{/each}
</div>
