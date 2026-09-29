{{template "layouts/header.tpl" .}}

<section>
    <h1>Event Explorer</h1>

    <p>Discover music and sports events in your city.</p>

    <form action="/events" method="GET">
        <label for="city">City</label>
        <input
            type="text"
            id="city"
            name="city"
            placeholder="Search for a city"
        >

        <input
            type="hidden"
            name="countryCode"
            value=""
        >

        <button type="submit">
            Search
        </button>
    </form>
</section>

{{template "layouts/footer.tpl" .}}