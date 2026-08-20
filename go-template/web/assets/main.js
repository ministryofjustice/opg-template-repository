import {initAll} from 'govuk-frontend';
import "govuk-frontend/dist/govuk/all.mjs";
import "opg-sirius-header/sirius-header.js";
import htmx from "htmx.org/dist/htmx.esm";

document.body.className += ' js-enabled' + ('noModule' in HTMLScriptElement.prototype ? ' govuk-frontend-supported' : '');
initAll();

window.htmx = htmx
htmx.logAll();
htmx.config.responseHandling = [{code:".*", swap: true}]

const formToggler = (suffix) => {
    return {
        resetAll: resetAll(suffix),
        show: show(suffix),
    }
}

const resetAll = (suffix) => () => {
    htmx.findAll(`[id$="-${suffix}"]`).forEach(element => {
        htmx.addClass(element, "hide");
        const input = element.querySelector("input");
        if (input) {
            input.setAttribute("disabled", "true");
            input.removeAttribute("max");
        }
    });
}

const show = (suffix) => (idName) => {
    document.querySelector(`#${idName}`).removeAttribute("disabled");
    htmx.removeClass(htmx.find(`#${idName}-${suffix}`), "hide")
}

// adding event listeners inside the onLoad function will ensure they are re-added to partial content when loaded back in
htmx.onLoad(() => {
    initAll();

    htmx.findAll(".moj-banner--success").forEach((element) => {
        element.addEventListener("click", () => {
            htmx.addClass(element, "hide");
            const url = new URL(window.location.href);
            if (url.searchParams.has('success')) {
                url.searchParams.delete('success');
                window.history.replaceState({}, '', url.toString());
            }
        });
    });
});