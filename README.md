# Gleeman

Gleeman is a toy static blog generator that aims to be a simple way to tell tales. The term "gleeman" is taken from Robert Jordan's epic fantasy series, _The Wheel of Time_. A gleeman is a traveling storyteller and entertainer.

## Quick Start

### Install
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

### Create a Blog Post

Create a new Markdown file in the `tales/entries/` directory. The filename can be whatever you want, but we suggest using the slug of the title of your post. For example, a blog post titled "My 5 Favorite Dogs" might have a filename of `my-five-favorite-dogs.md`. 

Write your blog post using Markdown.

### Build the Blog Post

Once your Markdown file is created, simply run `gleeman build` and it will generate the HTML for the site.

## Customizing your Blog

Update the following files to customize your blog:

* `_entry.html` - represents the template for your blog entry.
* `_main.html` - represents the main (index) page of your blog.

### Settings

The `settings.yaml` file contains only one required field at the moment, and that is `name`, which should be the name of your blog. By default, this is displayed in the `<title>` tags and at the top of your blog page.

Another option is `base_url`, if you want to define the location of this app. Normally, this is not needed, but if your files are being generated in the wrong place, setting this to the path of the gleeman executable this may resolve the issue.

### Wish List

* actually write blog posts