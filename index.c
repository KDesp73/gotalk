#include <webc/webc-core.h>
#include <webc/webc-actions.h>

char* Index() {
    char* buffer = NULL;

    WEBC_HtmlStart(&buffer, "en");
    WEBC_Head(&buffer, "gotalk", 
        META_AUTHOR_TAG("Konstantinos Despoinidis"),
        META_KEYWORDS_TAG("go, api, comments"),
        NULL
    );

    WEBC_StyleStart(&buffer, NO_ATTRIBUTES);
        WEBC_IntegrateFile(&buffer, "https://raw.githubusercontent.com/sindresorhus/github-markdown-css/main/github-markdown-light.css");
    WEBC_StyleEnd(&buffer);

    WEBC_BodyStart(&buffer, STYLE("padding: 0; margin: 0;"));
        WEBC_DivStart(&buffer, WEBC_UseModifier((Modifier){.class = "markdown-body", .style = "margin: 5% 25%;"}));
            WEBC_IntegrateFile(&buffer, "./README.md");
        WEBC_DivEnd(&buffer);
    WEBC_BodyEnd(&buffer);

    WEBC_HtmlEnd(&buffer);

    return buffer;
}

int main(int argc, char** argv) {
    WebcAction action = WEBC_ParseCliArgs(argc, argv);
    
    Tree tree = WEBC_MakeTree("docs", 
        WEBC_MakeRoute("/", Index()),
        NULL
    );

    WEBC_HandleAction(action, tree);
    return 0;
}
