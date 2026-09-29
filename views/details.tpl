{{template "layouts/header.tpl" .}}

<section>
    <h1>Event Details</h1>

    <p>
        Event ID:
        <strong>{{.EventID}}</strong>
    </p>

    <p>
        Event details will appear here.
    </p>

    <a href="/events">Back to events</a>
</section>

{{template "layouts/footer.tpl" .}}