FROM chromedp/headless-shell:92.0.4515.20

COPY files/entrypoint.sh /
COPY settings.ini /
COPY files/bgw210 /

ENTRYPOINT ["/entrypoint.sh"]


