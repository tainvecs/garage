import os
import json

from flask import Flask, request, Response, abort
import meilisearch as ms


class EndpointAction():

    def __init__(self, action):
        self.action = action

    def __call__(self, *args, **kwargs):
        return self.action(**kwargs)


class MeilisearchAPIServer():

    app = None
    ms_client = None
    ms_doc_index = None

    def __init__(self, name, host="0.0.0.0", port=5000, debug=False):

        # init flask app
        self.host = host
        self.port = port
        self.debug = debug
        self.app = Flask(name)

        # init meilisearch client
        ms_url = os.getenv("MEILISEARCH_URL")
        ms_doc_index = os.getenv("MEILISEARCH_DOCUMENT_INDEX")
        if not ms_url:
            self.app.logger.error("Missing env MEILISEARCH_URL while connecting to Meilisearch server.")
            abort(500)
        elif not ms_doc_index:
            self.app.logger.error("Missing env MEILISEARCH_DOCUMENT_INDEX while connecting to Meilisearch server.")
            abort(500)
        else:
            self.ms_client = ms.Client(ms_url)
            self.ms_doc_index = ms_doc_index

        # init API
        self.add_endpoint(endpoint="/v1/ms/docs", endpoint_name="search", handler=self.search, req_methods=["GET"])
        self.add_endpoint(endpoint="/v1/ms/docs", endpoint_name="index", handler=self.index, req_methods=["POST"])

    def run(self):
        self.app.run(debug=self.debug, port=self.port, host=self.host)

    def add_endpoint(self, endpoint=None, endpoint_name=None, handler=None, req_methods=["GET"]):
        self.app.add_url_rule(endpoint, endpoint_name, EndpointAction(handler), methods=req_methods)

    def index(self):

        # parse and check request
        req_data = request.get_json().get("data")
        if not req_data:
            self.app.logger.error("Missing index request data.")
            return Response(status=500, headers={})

        # run index
        resp = self.ms_client.index(self.ms_doc_index).add_documents(req_data)
        return Response(
            response=json.dumps(resp),
            status=200,
            headers={"Content-Type": "application/json"}
        )

    def search(self):

        # parse and check request
        query = request.args.get('query')
        if not query:
            self.app.logger.error("Missing query in search request.")
            return Response(status=400, headers={})

        # run search
        resp = self.ms_client.index(self.ms_doc_index).search(query)
        return Response(
            response=json.dumps(resp),
            status=200,
            headers={"Content-Type": "application/json"}
        )


if __name__ == "__main__":

    ms_api_server = MeilisearchAPIServer(
        name="meilisearch_api",
        host=os.getenv("FLASK_HOST"),
        port=os.getenv("FLASK_PORT"),
        debug=True
    )
    ms_api_server.run()
