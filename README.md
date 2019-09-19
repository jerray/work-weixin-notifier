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
