interface UploadParams {
    timestamp: number;
    signature: string;
    cloudName: string;
    apiKey: string;
    uploadPreset: string;
}

interface CloudinaryResponse {
    secure_url: string;
    public_id: string;
    asset_id: string;
}

export const uploadService = {
    async getUploadParams(): Promise<UploadParams> {
        const response = await fetch('https://api.flippy.toys/api/v1/toys/upload/params', {
            headers: {
                'X-Telegram-Init-Data': window.Telegram.WebApp.initData
            }
        });

        if (!response.ok) {
            throw new Error('Failed to get upload parameters');
        }

        return response.json();
    },

    async uploadImage(file: File, params: UploadParams): Promise<CloudinaryResponse> {
        const formData = new FormData();
        formData.append('file', file);
        formData.append('timestamp', params.timestamp.toString());
        formData.append('signature', params.signature);
        formData.append('api_key', params.apiKey);
        formData.append('upload_preset', params.uploadPreset);

        const response = await fetch(
            `https://api.cloudinary.com/v1_1/${params.cloudName}/image/upload`,
            {
                method: 'POST',
                body: formData
            }
        );

        if (!response.ok) {
            throw new Error('Failed to upload image');
        }

        return response.json();
    },

    async uploadMultiple(files: File[]): Promise<CloudinaryResponse[]> {
        const params = await this.getUploadParams();
        const uploadPromises = files.map(file => this.uploadImage(file, params));
        return Promise.all(uploadPromises);
    }
};