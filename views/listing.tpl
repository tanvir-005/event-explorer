{{template "layouts/header.tpl" .}}

<section>
    <h1>Events</h1>

    <p>
        Selected city:
        <strong>{{.City}}</strong>
        {{if .CountryCode}}
            <span>({{.CountryCode}})</span>
        {{end}}
    </p>
</section>

<section>
    <h2>Music</h2>

    <p>Music events will appear here.</p>
</section>

<section>
    <h2>Sports</h2>

    <p>Sports events will appear here.</p>
</section>

<a href="/">Change city</a>

{{template "layouts/footer.tpl" .}}