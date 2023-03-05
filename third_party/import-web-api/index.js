import http from 'http';
import { Form } from 'multiparty';
import { exec } from 'child_process';
import * as dotenv from 'dotenv'

dotenv.config();

const host = process.env.server_host;
const port = process.env.server_port;
const leaksDbFilePath = process.env.leaksdb_fp;

const corsHeaders = {
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Methods': 'OPTIONS, POST',
};

const leakFormSchema = {
    context: new String(),
    leakFile: new String(),
    shareDate: new Date(),
    leakers: new String(),
    platforms: new String(),
};

function triggerImportLeak(leakForm) {
    const cmd = `import --database-path="${leaksDbFilePath}" --leak-path="${leakForm.leakFile}" --context="${leakForm.context}" --platforms="${leakForm.platforms}" --share-date="${leakForm.shareDate}" --leakers="${leakForm.leakers}`;
    console.info(cmd);

    exec(cmd);
}

const server = http.createServer(function (req, res) {
    if (req.method === 'POST') {
        const reqForm = new Form();

        reqForm.parse(req, function (err, fields, files) {
            if (!err) {
                const leakForm = Object.assign({}, leakFormSchema);

                leakForm.context = `${fields.context[0]}`;
                leakForm.shareDate = new Date(Number.parseInt(`${fields.shareDateMS[0]}`)).toISOString();
                leakForm.platforms = `${fields.platforms[0]}`.split(',');
                leakForm.leakers = `${fields.leakers[0]}`.split(',');
                leakForm.leakFile = files.leakFile[0].path;

                console.info(`triggering leak import with the following data: ${leakForm}`);

                triggerImportLeak(leakForm);
            } else {
                console.error(err);
            }
        });

        res.writeHead(200, { 'Content-Type': 'text/plain', ...corsHeaders });
        res.end(null);
    } else {
        res.writeHead(405, { 'Content-Type': 'text/plain' });
        res.end('Method Not Allowed\n');
    }
});

server.listen(port, host);
console.log(`Server running on port ${port}`);