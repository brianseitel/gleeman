# Gleeman

Gleeman is a toy static blog generator that aims to be a simple way to tell tales. The term "gleeman" is taken from Robert Jordan's epic fantasy series, _The Wheel of Time_. A gleeman is a traveling storyteller and entertainer.

## Quick Start

To get started, run `gleeman init`. This will instantiate your blog skeleton. The directory structure looks like this:

```
myblog/
    public/
        assets/
            main.css
    tales/
        entries/
            sample.md
        layout/
            _layout.html
        settings.yaml
```

### Customizing your Blog

The two files used for customization are `main.css` for styling and `_layout.html` for the layout. 

The `main.css` file is a minified theme that you can use. Feel free to customize it as much or as little as you want.

The `_layouts.html` file contains the header, footer, and body of the blog. You can customize this if you want.

### Settings

The `settings.yaml` file contains only one required field at the moment, and that is `name`, which should be the name of your blog. By default, this is displayed in the `<title>` tags and at the top of your blog page.

Another option is `base_url`, if you want to define the location of this app. Normally, this is not needed, but if your files are being generated in the wrong place, setting this to the path of the gleeman executable this may resolve the issue.