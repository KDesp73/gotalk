#define CLIB_IMPLEMENTATION
#include <webc/extern/clib.h>
#define HTTPD_IMPLEMENTATION
#include <webc/extern/httpd.h>
#include <webc/webc-core.h>
#include <webc/webc-actions.h>

typedef enum {
    DARK, LIGHT
} Theme;

char* Error(size_t status, Theme theme) {
    char* buffer = NULL;

    WEBC_HtmlStart(&buffer, "en");
    WEBC_Head(&buffer, "gotalk", 
        META_AUTHOR_TAG("Konstantinos Despoinidis"),
        META_KEYWORDS_TAG("go, api, comments"),
        NULL
    );

    WEBC_StyleStart(&buffer, NO_ATTRIBUTES);
        if (theme == DARK)
            WEBC_IntegrateFile(&buffer, "https://raw.githubusercontent.com/sindresorhus/github-markdown-css/main/github-markdown-dark.css");
        else 
            WEBC_IntegrateFile(&buffer, "https://raw.githubusercontent.com/sindresorhus/github-markdown-css/main/github-markdown-light.css");
    WEBC_StyleEnd(&buffer);

    Modifier body_mod = {0};
    if (theme == DARK) {
        body_mod.style = strdup("padding: 0; margin: 0; background-color: #0d1117;");
    } else {
        body_mod.style = strdup("padding: 0; margin: 0; background-color: #ffffff;");
    }

    WEBC_BodyStart(&buffer, WEBC_UseModifier(body_mod));
        WEBC_DivStart(&buffer, WEBC_UseModifier((Modifier){.class = "markdown-body", .style = "margin: 5% 25%;"}));
            char* status_str = clib_format_text("Error %zu", status);
            WEBC_H1(&buffer, NO_ATTRIBUTES, status_str);
            WEBC_H3(&buffer, NO_ATTRIBUTES, status_message(status));
            free(status_str);
        WEBC_DivEnd(&buffer);
    WEBC_BodyEnd(&buffer);

    free((char*) body_mod.style);

    WEBC_HtmlEnd(&buffer);

    return buffer;
}

char* Index(Theme theme)
{
    char* buffer = NULL;

    WEBC_HtmlStart(&buffer, "en");
    WEBC_Head(&buffer, "gotalk", 
        META_AUTHOR_TAG("Konstantinos Despoinidis"),
        META_KEYWORDS_TAG("go, api, comments"),
        NULL
    );

    WEBC_StyleStart(&buffer, NO_ATTRIBUTES);
        if (theme == DARK)
            WEBC_IntegrateFile(&buffer, "https://raw.githubusercontent.com/sindresorhus/github-markdown-css/main/github-markdown-dark.css");
        else 
            WEBC_IntegrateFile(&buffer, "https://raw.githubusercontent.com/sindresorhus/github-markdown-css/main/github-markdown-light.css");
    WEBC_StyleEnd(&buffer);

    Modifier body_mod = {0};
    if (theme == DARK) {
        body_mod.style = strdup("padding: 0; margin: 0; background-color: #0d1117;");
    } else {
        body_mod.style = strdup("padding: 0; margin: 0; background-color: #ffffff;");
    }

    WEBC_BodyStart(&buffer, WEBC_UseModifier(body_mod));
        WEBC_DivStart(&buffer, WEBC_UseModifier((Modifier){.class = "markdown-body", .style = "margin: 5% 25%;"}));
            WEBC_IntegrateFile(&buffer, "./README.md");
        WEBC_DivEnd(&buffer);
    WEBC_BodyEnd(&buffer);

    free((char*) body_mod.style);

    WEBC_HtmlEnd(&buffer);

    return buffer;
}

int main(int argc, char** argv)
{
    WebcAction action = WEBC_ParseCliArgs(argc, argv);
    
    Tree tree = WEBC_MakeTree("docs", 
        WEBC_MakeRoute("/", Index(LIGHT)),
        WEBC_MakeRoute("/dark", Index(DARK)),
        WEBC_MakeRoute("/404", Error(404, LIGHT)),
        WEBC_MakeRoute("/404/dark", Error(404, DARK)),
        NULL
    );

    WEBC_HandleAction(action, tree);
    return 0;
}
