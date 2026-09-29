{{template "layouts/header.tpl" .}}

<section class="events-page">
    <div class="page-header">
        <div>
            <h1>Events in {{.City}}</h1>

            {{if .CountryCode}}
                <p>Country: {{.CountryCode}}</p>
            {{end}}
        </div>

        <a href="/">Change city</a>
    </div>

    {{if .Error}}
        <div class="message error">
            {{.Error}}
        </div>
    {{else}}

        <section class="event-section">
            <h2>Music</h2>

            {{if .MusicError}}
                <div class="message error">
                    Unable to load music events.
                </div>
            {{else if .MusicEvents}}
                <div class="event-grid">
                    {{range .MusicEvents}}
                        <article class="event-card">
                            {{if .ImageURL}}
                                <img
                                    src="{{.ImageURL}}"
                                    alt="{{.Name}}"
                                >
                            {{else}}
                                <div class="event-image-placeholder">
                                    No image available
                                </div>
                            {{end}}

                            <div class="event-card-content">
                                <h3>{{.Name}}</h3>

                                <p>{{.Date}}</p>

                                {{if .Venue}}
                                    <p>{{.Venue}}</p>
                                {{end}}

                                <a href="/events/{{.ID}}">
                                    View Details
                                </a>
                            </div>
                        </article>
                    {{end}}
                </div>
            {{else}}
                <p class="empty-message">
                    No music events found for this city.
                </p>
            {{end}}
        </section>

        <section class="event-section">
            <h2>Sports</h2>

            {{if .SportsError}}
                <div class="message error">
                    Unable to load sports events.
                </div>
            {{else if .SportsEvents}}
                <div class="event-grid">
                    {{range .SportsEvents}}
                        <article class="event-card">
                            {{if .ImageURL}}
                                <img
                                    src="{{.ImageURL}}"
                                    alt="{{.Name}}"
                                >
                            {{else}}
                                <div class="event-image-placeholder">
                                    No image available
                                </div>
                            {{end}}

                            <div class="event-card-content">
                                <h3>{{.Name}}</h3>

                                <p>{{.Date}}</p>

                                {{if .Venue}}
                                    <p>{{.Venue}}</p>
                                {{end}}

                                <a href="/events/{{.ID}}">
                                    View Details
                                </a>
                            </div>
                        </article>
                    {{end}}
                </div>
            {{else}}
                <p class="empty-message">
                    No sports events found for this city.
                </p>
            {{end}}
        </section>

    {{end}}
</section>

{{template "layouts/footer.tpl" .}}