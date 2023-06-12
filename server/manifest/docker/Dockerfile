FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
ADD hack                $WORKDIR/hack/
ADD manifest/config                $WORKDIR/manifest/config/
ADD resource                $WORKDIR/resource/
ADD ./temp/linux_amd64/hotgo $WORKDIR/hotgo
ADD ./manifest/docker/entrypoint.sh $WORKDIR/entrypoint.sh
RUN chmod +x $WORKDIR/hotgo
RUN chmod +x $WORKDIR/entrypoint.sh


###############################################################################
#                                   START
###############################################################################

WORKDIR $WORKDIR
CMD ./entrypoint.sh

