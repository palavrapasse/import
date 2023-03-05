import type { Form } from "@components";

const importWebApiUrl = 'http://0.0.0.0:55545';

export default function (form: Form): Promise<Response> {
    const formData = new FormData();

    formData.set('context', form.context);
    formData.set('shareDateMS', form.shareDateMS.toString());
    formData.set('platforms', form.platforms.join(','));
    formData.set('leakers', form.leakers.join(','));
    formData.set('leakFile', form.leakFile);

    return fetch(
        importWebApiUrl,
        {
            method: 'POST',
            body: formData,
        }
    );

}