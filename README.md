# Work Weixin Notifier

Send notification to Work Weixin group using Group Bot.

## Usage

```yaml
- uses: jerray/work-weixin-notifier@v1.0.0
  if: always()
  with:
    key: ${{ secrets.weixin_key }}
    type: text
    content: '{{ github.Repository }} build finished'
    status: ${{ job.status }}
```

Content is parsed and rendered with handlebars template parser. All default environments are
grouped under `github` object. Input values are grouped under `inputs` object. Keys are **CamelCased**.

Note: `GITHUB_SHA` in the template renamed to `github.Commit`.

### Build Context

The notifier provides some useful context values for rendering template. They are grouped under
`build` object.

* `build.Owner` is the repository owner, a user name or an organization name
* `build.Name` is repository's name
* `build.Branch` is the commit's branch name only if ref is branch ref
* `build.Tag` is the tag name only if ref is tag ref
* `build.Commit` is the shortcut of commit SHA
* `build.Link` is commit URL on GitHub if ref is branch, or tag tree URL if ref is tag

### Content Example

Here is a markdown type message content template. Set input `type` to `markdown` to use it.

```handlebars
# {{#equal inputs.Status "success"}}🔵{{/equal}}{{#equal inputs.Status "failure"}}🔴{{/equal}} {{ github.Repository }}
{{#if build.Tag}}[Tag {{ build.Tag }}]({{ build.Link }}){{else}}[Commit {{ build.Commit }}]({{ build.Link }}) on branch **{{ build.Branch }}**{{/if}} build {{ inputs.Status }}.
Click [here]({{ build.Link }}/checks) to check the detailed information.
```
